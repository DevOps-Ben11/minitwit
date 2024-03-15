import { useAuth } from '@/lib/hooks/useAuth'
import React from 'react'
import { Link } from 'react-router-dom'

type Props = {
  children: React.ReactNode
}

const Layout = ({ children }: Props) => {
  const { isAuthenticated, username } = useAuth()

  return (
    <div className='page'>
      <h1>MiniTwit</h1>

      <div className='navigation'>
        {isAuthenticated ? (
          <>
            <Link to='/timeline'>my timeline</Link> |
            <Link to='/public_timeline'>public timeline</Link> |
            <Link to='/logout'>sign out [{username}]</Link>
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
