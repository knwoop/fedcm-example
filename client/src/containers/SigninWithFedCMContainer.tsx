import { randomBytes } from 'crypto'
import Button from 'components/atoms/Button'
import { useAuthContext } from 'contexts/AuthContext'
import { useGlobalSpinnerActionsContext } from 'contexts/GlobalSpinnerContext'

interface SigninWithFedCMContainerProps {
  onSignin: (error?: Error) => void
}

/**
 * サインインフォームコンテナ
 */
const SigninWithFedCMContainer = ({
  onSignin,
}: SigninWithFedCMContainerProps) => {
  const { signinWithFedCM } = useAuthContext()
  const setGlobalSpinner = useGlobalSpinnerActionsContext()
  // サインインボタンを押された時のイベントハンドラ
  const handleSignin = async () => {
    const N = 16
    const nonce = randomBytes(N).toString('base64').substring(0, N)
    try {
      // ローディングスピナーを表示する
      setGlobalSpinner(true)
      await signinWithFedCM(nonce)
      onSignin && onSignin()
    } catch (err: unknown) {
      if (err instanceof Error) {
        // エラーの内容を表示
        window.alert(err.message)
        onSignin && onSignin(err)
      }
    } finally {
      setGlobalSpinner(false)
    }
  }

  return (
    <Button variant={'secondary'} width="100%" onClick={() => handleSignin()}>
      FedCM demo
    </Button>
  )
}

export default SigninWithFedCMContainer
