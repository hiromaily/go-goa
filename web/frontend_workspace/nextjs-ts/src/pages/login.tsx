import type { NextPage, GetStaticProps } from 'next'
import * as React from 'react'
import Layout from '../components/layout'
import SignIn from '../components/login'

type LoginProps = {
  message: string
}

const LoginPage: NextPage<LoginProps> = ({ message }) => {
  return (
    <Layout title='Top Page'>
      <SignIn message={message} />
    </Layout>
  )
}

export default LoginPage

// For SSG
export const getStaticProps: GetStaticProps<LoginProps> = async (context) => {
  const timestamp = new Date().toLocaleString()
  const message = `this page was rendered by calling getStaticProps() at ${timestamp}`
  return {
    props: {
      message,
    },
  }
}
