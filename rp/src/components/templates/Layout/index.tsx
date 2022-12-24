interface LayoutProps {
  children: React.ReactNode
}

const Layout = ({ children }: LayoutProps) => {
  return (
    <>
      <main>{children}</main>
    </>
  )
}

export default Layout
