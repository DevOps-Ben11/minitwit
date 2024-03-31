import axios from 'axios'
import { useEffect, useState } from 'react'
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

  useEffect(() => {
    const fetchMessages = async () => {
      const response = await axios.get<Response>('/api/timeline')
      console.log(response)
      setMessages(response.data.Messages)
      setUser(response.data.User)
    }

    fetchMessages()
  }, [])

  return (
    <div>
      <h2>My Timeline</h2>

      {user && (
        <div className='twitbox'>
          <h3>What's on your mind {user.Username}?</h3>
          <form>
            <p>
              <input type='text' name='text' />
              <input type='submit' value='Share' />
            </p>
          </form>
        </div>
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
    </div>
  )
}
