import type { NextPage, GetStaticProps } from 'next'
import * as React from 'react'
import Layout from '../components/layout'
import Top from '../components/top'
//import Style from '../styles/index.module.scss'

type IndexProps = {
  message: string
}

const Index: NextPage<IndexProps> = ({ message }) => {
  return (
    <Layout title='Top Page'>
      <Top message={message} />
    </Layout>
  )
}

export default Index

// For SSG
export const getStaticProps: GetStaticProps<IndexProps> = async (context) => {
  const timestamp = new Date().toLocaleString()
  const message = `this page was rendered by calling getStaticProps() at ${timestamp}`
  return {
    props: {
      message,
    },
  }
}
