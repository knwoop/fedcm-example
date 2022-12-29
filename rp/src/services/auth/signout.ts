import type { ApiContext } from 'types'
import { fetcher } from 'utils'

/**
 * Sing out API
 * @param context API context
 * @returns Sigin out
 */
const signout = async (context: ApiContext): Promise<{ message: string }> => {
  return await fetcher(
    `${context.apiRootUrl.replace(/\/$/g, '')}/auth/signout`,
    {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
    },
  )
}

export default signout
