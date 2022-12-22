import { useRouter } from 'next/router'
import Button from '../components/atoms/Button'
import Text from '../components/atoms/Text'
import Box from '../components/layout/Box'
import Flex from '../components/layout/Flex'
import Layout from '../components/templates/Layout'

export default function Home() {
  const router = useRouter()
  const toSigninOnClick = async () => {
    await router.push('/signin')
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
          <Box>
            <Text variant={'large'} color={'white'}>
              Home
            </Text>
          </Box>
          <Box width="100%" margin="10px">
            <Button
              variant={'secondary'}
              width="100%"
              onClick={() => toSigninOnClick()}
            >
              Home
            </Button>
          </Box>
        </Flex>
      </Flex>
    </Layout>
  )
}
