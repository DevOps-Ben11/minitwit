import { useLocalStorage } from 'usehooks-ts'

export const useAuth = () => {
  const [username, setUsername] = useLocalStorage<string | null>(
    'username',
    null
  )

  return {
    username,
    setUsername,
    isAuthenticated: Boolean(username),
  }
}
