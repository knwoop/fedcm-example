import { ApiContext } from 'types'

export type SigninWithFedCMParams = {
  nonce: string
}

/**
 * authentication API with FedCM
 * @param context API Context
 * @param params Parameter
 * @returns sign-in user
 */
const signinWithFedCM = async (
  context: ApiContext,
  params: SigninWithFedCMParams,
): Promise<Credential | null> => {
  return navigator.credentials.get({
    identity: {
      providers: [
        {
          configURL: process.env.NEXT_PUBLIC_FEDCM_CONFIG_URL,
          clientId: process.env.NEXT_PUBLIC_AUTH_CLIENT_ID,
          nonce: params.nance,
        },
      ],
    },
  })
}

export default signinWithFedCM
