import React from 'react';
import ProjectActions from '../../actions/projectActions';
import ProjectStore from '../../stores/projectStore';
import { Link, browserHistory } from 'react-router';
import styles from './style.css';

import TextField from 'material-ui/lib/text-field';
import RaisedButton from 'material-ui/lib/raised-button';
import FloatingActionButton from 'material-ui/lib/floating-action-button';
import ContentBack from 'material-ui/lib/svg-icons/navigation/arrow-back';


export default class CreateProjectPage extends React.Component {
  componentDidMount(){
      this.unsubscribe = ProjectStore.listen((state) => {this.onListChange(state);});
  }
  
  componentWillUnmount(){
      this.unsubscribe();
  }
  
  onListChange(state){
      if (state.refreshProject){
          browserHistory.push("/project/"+this.state.name);
      }
  }
  
  projectCreate(){
      ProjectActions.createProject(this.state.name, this.state.compose);
  }
  
  handleChange(event){
      var data = {};
      data[event.target.id] = event.target.value;
      this.setState(data);
  }
  
  render() {
      return (
        <div className={styles.content}>
            <TextField id="name" style={{ width: "100%"}} type="text" placeholder="name" onChange={(event) => {this.handleChange(event);}}/>
            <textarea id="compose" style={{width: "100%", height: "400px"}} placeholder="" onChange={(event) => {this.handleChange(event);}}/>
            <center><RaisedButton onClick={() => {this.projectCreate();}}>Create</RaisedButton></center>
            <FloatingActionButton mini={true} style={{marginTop: 20, marginBottom:20}} onClick={() => {browserHistory.push('/project');}}>
                <ContentBack />
            </FloatingActionButton>
        </div>
      ); 
  }
}
