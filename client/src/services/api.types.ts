export type UserTimelineResponse = {
  Followed: boolean
  Messages: Message[]
  User: User
  Profile: User
}

export type PublicTimelineResponse = {
  Messages: Message[]
}

export type TimelineResponse = {
  Messages: Message[]
  User: User
}

export type Message = {
  Username: string
  Pub_date: number
  Email: string
  Text: string
}

type User = {
  Username: string
  User_id: number
}
