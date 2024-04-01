import { PageWrapper } from '@/components/PageWrapper'
import axios from 'axios'
import { useEffect, useState } from 'react'
import Gravatar from 'react-gravatar'
import { Link, useNavigate, useParams } from 'react-router-dom'

type Response = {
  Followed: boolean
  Messages: Message[]
  User: User
  Profile: User
}

type Message = {
  Username: string
  Pub_date: number
  Email: string
  Text: string
}

type User = {
  Username: string
  User_id: number
}

export const Timeline = () => {
  const { username } = useParams<{ username: string }>()
  const [messages, setMessages] = useState<Message[]>([])
  const [user, setUser] = useState<User | null>(null)
  const [profile, setProfile] = useState<User | null>(null)
  const [following, setFollowing] = useState<boolean>(false)
  const [flashMessage, setFlashMessage] = useState<string | null>(null)
  const navigate = useNavigate()

  useEffect(() => {
    if (!username) {
      return
    }

    const fetchMessages = async () => {
      try {
        const response = await axios.get<Response>(`/api/timeline/${username}`)

        setMessages(response.data.Messages)
        setUser(response.data.User)
        setProfile(response.data.Profile)
        setFollowing(response.data.Followed)
      } catch (error) {
        navigate('/')
      }
    }

    fetchMessages()
  }, [username, navigate])

  const handleUnfollow = async () => {
    try {
      await axios.post(`/api/${username}/unfollow`)
      setFlashMessage(`You are no longer following "${username}"`)
      setFollowing(false)
    } catch (error) {
      console.error(error)
    }
  }

  const handleFollow = async () => {
    try {
      await axios.post(`/api/${username}/follow`)
      setFlashMessage(`You are now following "${username}"`)
      setFollowing(true)
    } catch (error) {
      console.error(error)
    }
  }

  const renderFollowStatus = () => {
    if (profile?.User_id === user?.User_id) {
      return 'This is you!'
    }

    if (following) {
      return (
        <>
          You are currently following this user.{' '}
          <button className='unfollow' onClick={handleUnfollow}>
            Unfollow user
          </button>
          .
        </>
      )
    }

    return (
      <>
        You are not yet following this user.{' '}
        <button className='unfollow' onClick={handleFollow}>
          Follow user
        </button>
        .
      </>
    )
  }

  return (
    <PageWrapper flashMessage={flashMessage}>
      <h2>{username}'s Timeline </h2>

      <div className='followstatus'>{renderFollowStatus()}</div>

      <ul className='messages'>
        {messages && messages.length > 0 ? (
          messages.map((message, index) => (
            <li key={index}>
              <Gravatar email={message.Email} size={48} />

              <p>
                <strong>
                  <Link to={`/${message.Username}`}>{message.Username}</Link>
                </strong>{' '}
                {message.Text}{' '}
                <small>
                  &mdash; {new Date(message.Pub_date * 1000).toLocaleString()}
                </small>
              </p>
            </li>
          ))
        ) : (
          <li>
            <em>There's no message so far.</em>
          </li>
        )}
      </ul>
    </PageWrapper>
  )
}
