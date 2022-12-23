import React, { useContext } from 'react'
import useSWR from 'swr'
import signinWithIDToken from '../../clients/auth/signin-with-idtoken'
import signin from 'clients/auth/signin'
import signinWithFedCM from 'clients/auth/signin-with-fedcm'
import signout from 'clients/auth/signout'
import type { ApiContext, User } from 'types'

type AuthContextType = {
  authUser?: User
  isLoading: boolean
  signin: (username: string, password: string) => Promise<void>
  signinWithFedCM: (nonce: string) => Promise<void>
  signout: () => Promise<void>
  mutate: (
    data?: User | Promise<User>,
    shouldRevalidate?: boolean,
  ) => Promise<User | undefined>
}

type AuthContextProviderProps = {
  context: ApiContext
  authUser?: User
}

const AuthContext = React.createContext<AuthContextType>({
  authUser: undefined,
  isLoading: false,
  signin: async () => Promise.resolve(),
  signinWithFedCM: async () => Promise.resolve(),
  signout: async () => Promise.resolve(),
  mutate: async () => Promise.resolve(undefined),
})

export const useAuthContext = (): AuthContextType =>
  useContext<AuthContextType>(AuthContext)

export const AuthContextProvider = ({
  context,
  authUser,
  children,
}: React.PropsWithChildren<AuthContextProviderProps>) => {
  const { data, error, mutate } = useSWR<User>(
    `${context.apiRootUrl.replace(/\/$/g, '')}/me`,
  )
  const isLoading = !data && !error

  const signinInternal = async (username: string, password: string) => {
    await signin(context, { username, password })
    await mutate()
  }

  const signinWithFedCMInternal = async (nonce: string) => {
    const credential = await signinWithFedCM(context, { nonce: nonce })
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    const { token } = credential

    console.log(token)
    await signinWithIDToken(context, { id_token: token })
    await mutate()
  }

  const signoutInternal = async () => {
    await signout(context)
    await mutate()
  }

  return (
    <AuthContext.Provider
      value={{
        authUser: data ?? authUser,
        isLoading,
        signin: signinInternal,
        signinWithFedCM: signinWithFedCMInternal,
        signout: signoutInternal,
        mutate,
      }}
    >
      {children}
    </AuthContext.Provider>
  )
}
