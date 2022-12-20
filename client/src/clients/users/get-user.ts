import type { ApiContext, User } from 'types'
import { fetcher } from 'utils'

/**
 * GET UserAPI
 * @param context APIコンテキスト
 * @returns ユーザー
 */
const getUser = async (context: ApiContext): Promise<User> => {
  /**
   // getUSer API
   // sample response
   {
    "id": 1,
    "username": user1,
    "email": "user1@example.com",
  }
   */
  return await fetcher(`${context.apiRootUrl.replace(/\/$/g, '')}/user`, {
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
    },
  })
}

export default getUser
