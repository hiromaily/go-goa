import * as React from 'react'
import Avatar from '@mui/material/Avatar'
import Box from '@mui/material/Box'
import Card from '@mui/material/Card'
import CardContent from '@mui/material/CardContent'
import CardHeader from '@mui/material/CardHeader'
import Container from '@mui/material/Container'
import List from '@mui/material/List'
import ListItem from '@mui/material/ListItem'
import Paper from '@mui/material/Paper'
import Typography from '@mui/material/Typography'
import Grid from '@mui/material/Unstable_Grid2'
import { styled } from '@mui/material/styles'
import { useDarkMode } from 'usehooks-ts'
import BaseCard from './resume/card'
import JobHeader from './resume/jobHeader'

type ResumeProps = {
  message: string
}

// https://mui.com/material-ui/react-grid2/

const Item = styled(Paper)(({ theme }) => ({
  backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
  ...theme.typography.body2,
  padding: theme.spacing(1),
  textAlign: 'center',
  color: theme.palette.text.secondary,
}))

const Resume = ({ message }: ResumeProps) => {
  const { isDarkMode } = useDarkMode()

  return (
    <Container>
      <Box sx={{ flexGrow: 1, minHeight: '100px', marginTop: '50px' }}>
        <CardHeader
          avatar={<Avatar alt='Hiroki Yasui' src='/hy.png' sx={{ width: 24, height: 24 }} />}
          title='Resume'
          titleTypographyProps={{ variant: 'h4' }}
        />
      </Box>
      <Grid container spacing={2}>
        <Grid xs={3}>
          <BaseCard>
            <CardHeader titleTypographyProps={{ variant: 'subtitle1' }} title='Like' />
            <CardContent>
              <Item>
                <List>
                  <ListItem divider>Golang</ListItem>
                  <ListItem divider>Golang</ListItem>
                  <ListItem divider>Golang</ListItem>
                  <ListItem divider>Golang</ListItem>
                </List>
              </Item>
            </CardContent>
          </BaseCard>

          <Card>
            <CardHeader titleTypographyProps={{ variant: 'subtitle1' }} title='Dislike' />
            <CardContent>
              <Item>
                <List>
                  <ListItem divider>Golang</ListItem>
                  <ListItem divider>Golang</ListItem>
                  <ListItem divider>Golang</ListItem>
                  <ListItem divider>Golang</ListItem>
                </List>
              </Item>
            </CardContent>
          </Card>
        </Grid>

        <Grid xs={9}>
          <BaseCard>
            <JobHeader isDarkMode={isDarkMode} title='Blockchain Engineer at Datachain' />
            <CardContent>
              <Typography variant='subtitle1'>2021 July - At Present</Typography>
              <ul>
                <li>Developing various PoC</li>
                <li>Researching something</li>
              </ul>
            </CardContent>
          </BaseCard>

          <BaseCard>
            <JobHeader isDarkMode={isDarkMode} title='Blockchain Engineer at Datachain' />
            <CardContent>
              <Typography variant='subtitle1'>2021 July - At Present</Typography>
              <ul>
                <li>Developing various PoC</li>
                <li>Researching something</li>
              </ul>
            </CardContent>
          </BaseCard>
        </Grid>
      </Grid>
    </Container>
  )
}

export default Resume
