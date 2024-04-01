import { PageWrapper } from '@/components/PageWrapper'
import { TweetBox } from '@/components/TweetBox'
import axios from 'axios'
import { useCallback, useEffect, useState } from 'react'
import Gravatar from 'react-gravatar'
import { Link } from 'react-router-dom'

type Response = {
  Followed: boolean
  Messages: Message[]
  User: User
}

type Message = {
  Username: string
  Pub_date: number
  Email: string
  Text: string
}

type User = {
  Username: string
}

export const MyTimeline = () => {
  const [messages, setMessages] = useState<Message[]>([])
  const [user, setUser] = useState<User | null>(null)
  const [flashMessage, setFlashMessage] = useState<string | null>(null)

  const fetchMessages = useCallback(async () => {
    const response = await axios.get<Response>('/api/timeline')

    setMessages(response.data.Messages)
    setUser(response.data.User)
  }, [])

  useEffect(() => {
    fetchMessages()
  }, [fetchMessages])

  const handleAddMessageSuccess = () => {
    fetchMessages()
    setFlashMessage('Your message was recorded')
  }

  return (
    <PageWrapper flashMessage={flashMessage}>
      <h2>My Timeline</h2>

      {user && (
        <TweetBox
          handleCTA={handleAddMessageSuccess}
          username={user.Username}
        />
      )}

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
