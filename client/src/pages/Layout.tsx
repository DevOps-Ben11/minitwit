import { useAuth } from '@/lib/hooks/useAuth'
import axios from 'axios'
import React, { useState } from 'react'
import { Link } from 'react-router-dom'

type Props = {
  children: React.ReactNode
}

// function deleteCookie(name: string) {
//   document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
// }

const Layout = ({ children }: Props) => {
  const { isAuthenticated, username } = useAuth()
  const [error, setError] = useState<string | null>(null)

  const handleLogout = async (data: any) => {
    setError(null)
    try {
      await axios.get('/api/logout', data)

    } catch (error) {
      if (axios.isAxiosError(error) && error.response?.data?.error_msg) {
        setError(error.response.data.error_msg)
      }
    }
  };

  return (
    <div className='page'>
      <h1>MiniTwit</h1>

      <div className='navigation'>
        {isAuthenticated ? (
          <>
            <Link to='/timeline'>my timeline</Link> |
            <Link to='/public_timeline'>public timeline</Link> |
            <Link to='/logout' onClick={handleLogout}>sign out [{username}]</Link>
            {error && (
        <div className='error'>
          <strong>Error:</strong> {error}
        </div>
      )}
          </>
        ) : (
          <>
            <Link to='/public_timeline'>public timeline</Link> |
            <Link to='/register'>sign up</Link> |
            <Link to='/login'>sign in</Link>
          </>
        )}
      </div>

      <div className='body'>{children}</div>

      <div className='footer'>MiniTwit &mdash; A Go Application</div>
    </div>
  )
}

export default Layout