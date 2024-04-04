import { Message } from '@/services/api.types'
import Gravatar from 'react-gravatar'
import { Link } from 'react-router-dom'

type Props = {
  messages: Message[] | undefined
}

export const MessageList = ({ messages }: Props) => (
  <ul className='messages'>
    {messages && messages.length > 0 ? (
      messages.map((message, index) => (
        <li key={index}>
          <Gravatar email={message.Email} size={48} />

          <p>
            <strong>
              <Link to={`/timeline/${message.Username}`}>
                {message.Username}
              </Link>
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
)
