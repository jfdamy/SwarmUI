import React from 'react';
import ProjectActions from '../../actions/projectActions';
import ProjectStore from '../../stores/projectStore';
import { Link, browserHistory } from 'react-router';
import styles from './style.css';

import TextField from 'material-ui/lib/text-field';
import RaisedButton from 'material-ui/lib/raised-button';
import FloatingActionButton from 'material-ui/lib/floating-action-button';
import ContentBack from 'material-ui/lib/svg-icons/navigation/arrow-back';
import Table from 'material-ui/lib/table/table';
import TableHeaderColumn from 'material-ui/lib/table/table-header-column';
import TableRow from 'material-ui/lib/table/table-row';
import TableHeader from 'material-ui/lib/table/table-header';
import TableRowColumn from 'material-ui/lib/table/table-row-column';
import TableBody from 'material-ui/lib/table/table-body';
import Tabs from 'material-ui/lib/tabs/tabs';
import Tab from 'material-ui/lib/tabs/tab';
import Toolbar from 'material-ui/lib/toolbar/toolbar';
import ToolbarSeparator from 'material-ui/lib/toolbar/toolbar-separator';
import ToolbarGroup from 'material-ui/lib/toolbar/toolbar-group';


export default class ProjectPage extends React.Component {
    
  componentDidMount(){
      this.unsubscribe = ProjectStore.listen((state) => {this.onListChange(state);});
      ProjectActions.projectInfo(this.props.params.projectId);
  }
  
  componentWillUnmount(){
      this.unsubscribe();
  }
  
  onListChange(state){
      if (state.refreshProject){
          ProjectActions.projectInfo(this.props.params.projectId);
      } else if (state.returnToProjects) {
          browserHistory.push('/project');
      } else {
        this.setState(state);
      }
  }
  
  projectUp(){
      ProjectActions.projectUp(this.props.params.projectId);
  }
  
  projectStop(){
      ProjectActions.projectStop(this.props.params.projectId);
  }
  
  projectKill(){
      ProjectActions.projectKill(this.props.params.projectId);
  }
  
  projectDelete(){
      ProjectActions.projectDelete(this.props.params.projectId);
  }
  
  projectRemove(){
      ProjectActions.projectRemove(this.props.params.projectId);
  }
  
  projectScale(){
      ProjectActions.projectScale(this.props.params.projectId, this.state.services);
  }
  
  printPorts(ports){
      let ret = "";
      if(ports){
          ports.forEach( (value) => {
              ret += ' '+ value.PortHost + ' => '+ value.PortCont;
          });
      }
      return ret;
  }
  
  
  handleChangeField(event){
      var data = {};
      data[event.target.id] = event.target.value;
      this.setState(data);
  }
  
  handleChangeScale(event){
      var data = this.state.services ? this.state.services : [];
      data.push({
         ServiceName: event.target.id,
         Number: parseInt(event.target.value)
      });
      this.setState({services : data});
  }
  
  render() {
      var services = [];
      var project = null;
      if(this.state && this.state.project){
          services = this.state.project.Services;
          project = JSON.stringify(this.state.project, null, 2);
      }
      return (
        <div>
            <FloatingActionButton mini={true} style={{marginTop: 10}} 
                onClick={() => {browserHistory.push('/project');}}>
                <ContentBack />
            </FloatingActionButton>
            <center><h1 style={{marginBottom: 30}}>{this.props.params.projectId}</h1></center>
             <Toolbar>   
                <ToolbarGroup firstChild={true} float="left">
                    <RaisedButton onClick={() => {this.projectUp();}}>Up</RaisedButton>
                    <RaisedButton onClick={() => {this.projectDelete();}}>Delete</RaisedButton>
                </ToolbarGroup>
                
                <ToolbarGroup float="right"> 
                    <RaisedButton onClick={() => {this.projectStop();}}>Stop</RaisedButton>
                    <ToolbarSeparator />
                    <RaisedButton onClick={() => {this.projectKill();}}>Kill</RaisedButton>
                    <ToolbarSeparator />
                    <RaisedButton onClick={() => {this.projectRemove();}}>Remove</RaisedButton>
                    <ToolbarSeparator />
                    <RaisedButton onClick={() => {this.projectScale();}}>Scale</RaisedButton>
                </ToolbarGroup>
            </Toolbar>
            
                <div>
                {services.map(value => (
                    <div key={value.ServiceName}>
                        <h2>{value.ServiceName}</h2>
                        <div>
                            Scale : &nbsp;
                            <input type="number" id={value.ServiceName} onChange={ (event) => {this.handleChangeScale(event)}}></input>
                        </div>
                        <br />
                        <Table>
                            <TableHeader>
                            <TableRow>
                                <TableHeaderColumn>Name</TableHeaderColumn>
                                <TableHeaderColumn>isRunning</TableHeaderColumn>
                                <TableHeaderColumn>Port</TableHeaderColumn>
                            </TableRow>
                            </TableHeader>
                            <TableBody>
                                {value.Containers ? value.Containers.map(value => (
                                    <TableRow key={value.ContainerName}>
                                        <TableRowColumn>{value.ContainerName}</TableRowColumn>
                                        <TableRowColumn>{value.IsRunning ? "true" : "false"}</TableRowColumn>
                                        <TableRowColumn>{this.printPorts(value.Port)}</TableRowColumn>
                                    </TableRow>
                                )) : ""}
                            </TableBody>
                        </Table>
                    </div>
                ))}
            </div>
        </div>
      ); 
  }
}
