import { randomBytes } from 'crypto'
import { useEffect, useState } from 'react'
import Button from 'components/atoms/Button'
import { useAuthContext } from 'contexts/AuthContext'
import { useGlobalSpinnerActionsContext } from 'contexts/GlobalSpinnerContext'

interface SigninWithFedCMContainerProps {
  onSignin: (error?: Error) => void
}

/**
 * Sign-in with FedCM container
 */
const SigninWithFedCMContainer = ({
  onSignin,
}: SigninWithFedCMContainerProps) => {
  const { signinWithFedCM, isFedCMEnabled } = useAuthContext()

  const [mounted, setMounted] = useState(false)

  const setGlobalSpinner = useGlobalSpinnerActionsContext()
  // サインインボタンを押された時のイベントハンドラ
  const handleSignin = async () => {
    const N = 16
    const nonce = randomBytes(N).toString('base64').substring(0, N)
    try {
      setGlobalSpinner(true)
      await signinWithFedCM(nonce)
      onSignin && onSignin()
    } catch (err: unknown) {
      if (err instanceof Error) {
        window.alert(err.message)
        onSignin && onSignin(err)
      }
    } finally {
      setGlobalSpinner(false)
    }
  }

  useEffect(() => {
    if (isFedCMEnabled()) {
      setMounted(true)
    }
  }, [isFedCMEnabled])

  if (!mounted) {
    return null
  }

  return (
    <Button variant={'secondary'} width="100%" onClick={() => handleSignin()}>
      FedCM demo
    </Button>
  )
}

export default SigninWithFedCMContainer
