import type { NextPage } from 'next'
import { useRouter } from 'next/router'
import Button from '../components/atoms/Button'
import Text from '../components/atoms/Text'
import Box from 'components/layout/Box'
import Flex from 'components/layout/Flex'
import Layout from 'components/templates/Layout'

const PrivacyPolicyPage: NextPage = () => {
  const router = useRouter()
  const backHomeOnClick = async () => {
    await router.push('/')
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
            <Text color={'white'}>privacy policy...</Text>
          </Box>
          <Box width="100%" margin="10px">
            <Button
              variant={'secondary'}
              width="100%"
              onClick={() => backHomeOnClick()}
            >
              Home
            </Button>
          </Box>
        </Flex>
      </Flex>
    </Layout>
  )
}

export default PrivacyPolicyPage
