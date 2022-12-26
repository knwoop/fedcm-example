import { ApiContext } from 'types'

export type SigninWithFedCMParams = {
  configURL: string
  clientId: string
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
  if (typeof window === 'undefined') {
    // can't use on server side
    return null
  }
  if (!isFedCMEnabled()) {
    return null
  }

  return navigator.credentials.get({
    identity: {
      providers: [
        {
          configURL: params.configURL,
          clientId: params.clientId,
          nonce: params.nonce,
        },
      ],
    },
  })
}

export const isFedCMEnabled = (): boolean => !!window.IdentityCredential

export default signinWithFedCM
