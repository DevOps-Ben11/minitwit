import { useCookies } from 'react-cookie'
import { useLocalStorage } from 'usehooks-ts'

export const useAuth = () => {
  const [username, setUsername] = useLocalStorage<string | null>(
    'username',
    null
  )
  const [auth] = useCookies(['auth'])

  return {
    username,
    setUsername,
    isAuthenticated: Boolean(auth.auth),
  }
}
