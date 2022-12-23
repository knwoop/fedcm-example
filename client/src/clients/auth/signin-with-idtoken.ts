import { ApiContext, User } from 'types'
import { fetcher } from 'utils'

export type SigninWithIDTokenParams = {
  id_token: string
}

/**
 * authentication API (sign in)
 * @param context API Context
 * @param params Parameter
 * @returns sign-in user
 */
const signinWithIDToken = async (
  context: ApiContext,
  params: SigninWithIDTokenParams,
): Promise<User> => {
  return await fetcher(
    `${context.apiRootUrl.replace(/\/$/g, '')}/auth/idtoken`,
    {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(params),
    },
  )
}

export default signinWithIDToken
