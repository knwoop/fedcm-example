import { randomBytes } from 'crypto'
import type { NextPage } from 'next'
import { useRouter } from 'next/router'
import Button from '../components/atoms/Button'
import AppLogo from 'components/atoms/AppLogo'
import Box from 'components/layout/Box'
import Flex from 'components/layout/Flex'
import Layout from 'components/templates/Layout'
import SigninFormContainer from 'containers/SigninFormContainer'

const SigninPage: NextPage = () => {
  const router = useRouter()
  const handleSignin = async (err?: Error) => {
    if (!err) {
      const redurectTo = (router.query['redirect_to'] as string) ?? '/'

      console.log('Redirecting', redurectTo)
      await router.push(redurectTo)
    }
  }

  let nonce = ''
  const onSignInWithFedCMClick = async (err?: Error) => {
    const N = 16
    nonce = randomBytes(N).toString('base64').substring(0, N)
    const credential = await navigator.credentials.get({
      identity: {
        providers: [
          {
            configURL: process.env.NEXT_PUBLIC_FEDCM_CONFIG_URL,
            clientId: process.env.NEXT_PUBLIC_AUTH_CLIENT_ID,
            nonce: nonce,
          },
        ],
      },
    })
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    const { token } = credential
    console.log(token)
  }

  return (
    <Layout>
      <Flex
        paddingTop={2}
        paddingBottom={2}
        paddingLeft={{ base: 2, md: 0 }}
        paddingRight={{ base: 2, md: 0 }}
        justifyContent="center"
      >
        <Flex
          width="400px"
          flexDirection="column"
          justifyContent="center"
          alignItems="center"
        >
          <Box marginBottom={2}>
            <AppLogo />
          </Box>
          <Box width="100%">
            <SigninFormContainer onSignin={handleSignin} />
          </Box>
          <Box width="100%" margin="10px">
            <Button
              variant={'secondary'}
              width="100%"
              onClick={() => onSignInWithFedCMClick()}
            >
              FedCM demo
            </Button>
          </Box>
        </Flex>
      </Flex>
    </Layout>
  )
}

export default SigninPage
