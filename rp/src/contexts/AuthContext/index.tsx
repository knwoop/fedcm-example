import React, { useContext } from 'react'
import useSWR from 'swr'
import signinWithIDToken from '../../services/auth/signin-with-idtoken'
import signin from 'services/auth/signin'
import signinWithFedCM, {
  isFedCMEnabled,
} from 'services/auth/signin-with-fedcm'
import signout from 'services/auth/signout'
import type { ApiContext, User } from 'types'

type AuthContextType = {
  authUser?: User
  isLoading: boolean
  signin: (username: string, password: string) => Promise<void>
  signinWithFedCM: (nonce: string) => Promise<void>
  isFedCMEnabled: () => boolean
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
  isFedCMEnabled: () => false,
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
    const credential = await signinWithFedCM(context, {
      configURL: process.env.NEXT_PUBLIC_FEDCM_CONFIG_URL,
      clientId: process.env.NEXT_PUBLIC_AUTH_CLIENT_ID,
      nonce: nonce,
    })

    let token
    if (credential === null) {
      token = ''
    } else {
      token = credential.token
    }

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
        isFedCMEnabled: isFedCMEnabled,
        signout: signoutInternal,
        mutate,
      }}
    >
      {children}
    </AuthContext.Provider>
  )
}
