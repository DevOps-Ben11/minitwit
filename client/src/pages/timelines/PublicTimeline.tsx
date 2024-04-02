import { MessageList } from '@/components/MessageList'
import { PageWrapper } from '@/components/PageWrapper'
import { getPublicTimeline } from '@/services/api'
import { Message } from '@/services/api.types'
import { useEffect, useState } from 'react'
import { useLocation } from 'react-router-dom'

export const Public = () => {
  const { state } = useLocation()
  const [messages, setMessages] = useState<Message[]>([])

  useEffect(() => {
    const fetchMessages = async () => {
      const response = await getPublicTimeline()
      setMessages(response.data.Messages)
    }

    fetchMessages()
  }, [])

  return (
    <PageWrapper flashMessage={state && state.flashMessage}>
      <h2>Public Timeline</h2>

      <MessageList messages={messages} />
    </PageWrapper>
  )
}
