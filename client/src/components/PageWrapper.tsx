type Props = {
  flashMessage?: string | null
  children: React.ReactNode
}

export const PageWrapper = ({ flashMessage, children }: Props) => (
  <>
    {flashMessage && (
      <ul className='flashes'>
        <li>{flashMessage}</li>
      </ul>
    )}

    <div className='body'>{children}</div>
  </>
)
