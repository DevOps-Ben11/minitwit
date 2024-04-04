type Props = {
  flashMessage?: string | null
  children: React.ReactNode
}

export const PageWrapper = ({ flashMessage, children, ...props }: Props) => (
  <>
    {flashMessage && (
      <ul className='flashes'>
        <li>{flashMessage}</li>
      </ul>
    )}

    <div className='body' {...props}>
      {children}
    </div>
  </>
)
