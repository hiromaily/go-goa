import type { NextPage, GetStaticProps } from 'next'
import * as React from 'react'
import Resume from '../components/resume'

type ResumeProps = {
  message: string
}

const ResumePage: NextPage<ResumeProps> = ({ message }) => {
  return <Resume message={message} />
}

export default ResumePage

// For SSG
export const getStaticProps: GetStaticProps<ResumeProps> = async (context) => {
  const timestamp = new Date().toLocaleString()
  const message = `this page was rendered by calling getStaticProps() at ${timestamp}`
  return {
    props: {
      message,
    },
  }
}
