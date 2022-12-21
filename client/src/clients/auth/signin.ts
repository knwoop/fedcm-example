import { ApiContext, User } from 'types'
import { fetcher } from 'utils'

export type SigninParams = {
  username: string
  password: string
}

/**
 * authentication API (sign in)
 * @param context API Context
 * @param params Parameter
 * @returns sign-in user
 */
const signin = async (
  context: ApiContext,
  params: SigninParams,
): Promise<User> => {
  return await fetcher(
    `${context.apiRootUrl.replace(/\/$/g, '')}/auth/signin`,
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

export default signin
