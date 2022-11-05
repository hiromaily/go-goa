import * as React from 'react'
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
import Avatar from '@mui/material/Avatar'


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
  return (
    <Container>
      <Box sx={{ flexGrow: 1, minHeight: '100px', marginTop: '50px' }}>
        <CardHeader
          avatar={
            <Avatar
              alt="Hiroki Yasui"
              src="/hy.png"
              sx={{ width: 24, height: 24 }}
            />
          }
          title="Resume"
          titleTypographyProps={{ variant: 'h5' }}
        />
      </Box>
      <Grid container spacing={2}>
        <Grid xs={3}>
          <Card
            sx={{
              marginBottom: '30px',
            }}
          >
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
          </Card>

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
          <Item>1</Item>
        </Grid>
      </Grid>
    </Container>
  )
}

export default Resume
