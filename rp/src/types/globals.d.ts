interface CredentialRequestOptions {
  identity: {
    providers: {
      configURL: string
      clientId: string
      nonce: string
    }[]
  }
}

interface Credential {
  token: string
}

interface Window {
  IdentityCredential?: boolean
}

declare namespace NodeJS {
  interface ProcessEnv {
    readonly NODE_ENV: 'development' | 'production' | 'test'
    readonly NEXT_PUBLIC_FEDCM_CONFIG_URL: string
    readonly NEXT_PUBLIC_AUTH_CLIENT_ID: string

    readonly AAA: string
  }
}
