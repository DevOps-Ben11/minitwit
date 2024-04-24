import { PageWrapper } from '@/components/PageWrapper'
import { useAuth } from '@/lib/hooks/useAuth'
import { axios } from '@/services/api'
import { useEffect } from 'react'
import { toast } from 'react-toastify'

export const Logout = () => {
  const { setUsername } = useAuth()

  useEffect(() => {
    const signOut = async () => {
      try {
        await axios.post('/api/logout')
      } catch (error) {
        console.error(error)
      } finally {
        setUsername(null)
        toast.success('You were logged out')
      }
    }

    signOut()
  }, [setUsername])

  return <PageWrapper>Signing out...</PageWrapper>
}
