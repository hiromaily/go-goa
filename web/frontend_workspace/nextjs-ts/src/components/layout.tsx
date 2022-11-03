import { ReactNode } from 'react'
import Footer from './footer'
import Header from './header'
//import { MetaMaskProvider } from 'metamask-react'

type LayoutProps = {
  children?: ReactNode
  title: string
}

const Layout = ({ children, title }: LayoutProps) => {
  return (
    <>
      <Header title={title} />
      <main>{children}</main>
      <Footer />
    </>
  )
}

export default Layout
