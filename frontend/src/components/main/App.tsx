
import React from 'react';
import './index.css';
import imgAvatar from '../assets/logo/imgAvatar.png';
import {
  Link,
  BrowserRouter as Router,
  Route,
} from 'react-router-dom';
import SearchBar from '../search-bar/SearchBar';
import TaskContainer from '../task-container/TaskContainer';
import UploadTask from '../upload-task/UploadTask';
import '@patternfly/react-core/dist/styles/base.css';
import logo from '../assets/logo/main.png';
import Community from '../community/Community';
import Filter from '../filter/Filter';
import UserProfile from '../user-profile/UserProfile';
import {
  Button,
  ButtonVariant,
  ToolbarItem,
  Page,
  Brand,
  PageHeader,
  PageSection,
  Toolbar,
  ToolbarGroup,
  Avatar,
  Grid,
  GridItem,
} from '@patternfly/react-core';
import Detail from '../detail/Detail';
import BasicDetailParent from '../basic-detail/BasicDetailParent';
import BackgroundImageHeader from '../background-image/BackgroundImage';
import Login from '../Authentication/Login';
import SignupForm from '../Authentication/Signup';
import Footer from '../footer/Footer';
interface mainProps {

}
interface mainState {
  value: string;
}

const App: React.FC<mainProps> = () => {
  const logoProps = {
    href: '/',
    // eslint-disable-next-line no-console
    onClick: () => console.log('clicked logo'),
    target: '',
  };
  const logoutUser=()=>{
    localStorage.removeItem('token');
    localStorage.removeItem('usetrID');
    window.location.assign('/');
  };
  let userimage:any;
  let displayUpload:any ='';
  let authenticationButton;
  if (localStorage.getItem('token')===null) {
    authenticationButton= <Link to="/login">
      <span style={{marginRight: '1em', color: 'white'}}> Login </span>
    </Link>;
    displayUpload='';
  } else {
    authenticationButton= <Link to="/">
      <span style={{marginRight: '1em', color: 'white'}}
        onClick={logoutUser}> Logout </span>
    </Link>;
    displayUpload= <Link to="/upload">
      <span style={{marginLeft:
        '1.3em', marginRight:
         '-0.6em', color: 'white'}}>
    Add
      </span>
      {' '}
    </Link>;
    userimage = <Link to = "/user-profile">
      <Avatar src={imgAvatar} alt="" />
    </Link>;
  }

  // code for header contents
  const PageToolbar = (
    // eslint-disable-next-line react/jsx-filename-extension
    <div>
      <Toolbar>
        <ToolbarGroup>
          <ToolbarItem style={{color: 'white'}}>
            <Link to="/">
              <span style={{marginRight: '2em', color: 'white'}}>Home</span>
            </Link>

            <Link to="/">
              <span style={{marginRight: '2em', color: 'white'}}>Search</span>
            </Link>

            <Link to="/">
              <span style={{marginRight: '0.2em', color: 'white'}}>
                Community
              </span>
            </Link>
            {displayUpload}
            <Button id="default-example-uid-01"
              aria-label="Notifications actions"
              variant={ButtonVariant.plain}>
            </Button>
          </ToolbarItem>
          <ToolbarItem>
            {
              authenticationButton
            }


          </ToolbarItem>
          <ToolbarItem>
            <Link to="/"> {userimage}  </Link>
          </ToolbarItem>
        </ToolbarGroup>
      </Toolbar>
    </div>

  );
  const Header = (
    <PageHeader
      logo={<Brand src={logo} alt="Tekton Hub Logo" />}
      logoProps={logoProps}
      toolbar={PageToolbar}
    />
  );

  return (
    <Router>
      <Page header={Header}>
        <Route exact path="/" component={BackgroundImageHeader} />
        <Route exact path="/search" component={BackgroundImageHeader} />

        {/* </PageSection> */}
        <PageSection>
          <GridItem span={1}>

          </GridItem>
          <GridItem span={9}>
            <Route exact path="/" component={SearchBar} />
            <Route exact path="/detail/:taskId" component={BasicDetailParent} />
            <Route exact path="/search" component={SearchBar} />
          </GridItem>
        </PageSection>
        <PageSection>
          <Grid sm={6} md={4} lg={4} xl2={1} >
            <GridItem span={1} >
              <Route exact path="/" component={Filter} />
            </GridItem>
            <GridItem span={9} style = {{marginLeft: '2em'}} >
              <Route exact path="/user-profile" component={UserProfile}/>
              <Route exact path="/upload" component={UploadTask} />
              <Route exact path="/" component={TaskContainer} />
              <Route path="/search" component={TaskContainer} />
            </GridItem>
          </Grid>
          <Route exact path="/detail/:taskId" component={Detail} />
          <Route path="/community" component={Community} />
        </PageSection>
        <PageSection>
          <Route path='/login' component={Login}/>
          <Route path='/logout' component={Login}/>
          <Route path='/signup' component={SignupForm}/>
        </PageSection>

        <Footer />

      </Page>
    </Router>
  );
};

export default App;
