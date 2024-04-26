import { useAuth } from '@/lib/hooks/useAuth'
import { Link, useNavigate } from 'react-router-dom'

type Props = {
  children: React.ReactNode
}

const Layout = ({ children }: Props) => {
  const { isAuthenticated, username } = useAuth()
  const navigate = useNavigate()

  const handleLogout = async () => {
    try {
      navigate('/logout')
    } catch (error) {
      console.error(error)
    }
  }

  return (
    <div className='page'>
      <h1>MiniTwit</h1>

      <div className='navigation'>
        {isAuthenticated ? (
          <>
            <Link data-testid='my-timeline-link' to='/'>
              my timeline
            </Link>{' '}
            |
            <Link data-testid='public-timeline-link' to='/public'>
              public timeline
            </Link>{' '}
            |
            <button data-testid='sign-out' onClick={handleLogout}>
              sign out [{username}]
            </button>
          </>
        ) : (
          <>
            <Link to='/public'>public timeline</Link> |
            <Link to='/register'>sign up</Link> |
            <Link to='/login'>sign in</Link>
          </>
        )}
      </div>

      {children}

      <div className='footer'>MiniTwit v2 &mdash; A Go - React Application</div>
    </div>
  )
}

export default Layout
