type Props = {
  children: React.ReactNode
}

export const PageWrapper = ({ children, ...props }: Props) => (
  <>
    <div className='body' {...props}>
      {children}
    </div>
  </>
)
