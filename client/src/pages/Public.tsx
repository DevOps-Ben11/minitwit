import { PageWrapper } from '@/components/PageWrapper'
import axios from 'axios'
import { useEffect, useState } from 'react'
import Gravatar from 'react-gravatar'
import { Link } from 'react-router-dom'

type Response = {
  Followed: boolean
  Messages: Message[]
}

type Message = {
  Username: string
  Pub_date: number
  Email: string
  Text: string
}

export const Public = () => {
  const [messages, setMessages] = useState<Message[]>([])

  useEffect(() => {
    console.log('ASDASD')
    const fetchMessages = async () => {
      const response = await axios.get<Response>('/api/public')
      setMessages(response.data.Messages)
    }

    fetchMessages()
  }, [])

  return (
    <PageWrapper>
      <h2>Public Timeline</h2>

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
