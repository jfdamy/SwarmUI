import React from 'react';
import _ from 'lodash';
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
  
  autoscaling(type, serviceName, event){
      if (event.target.checked){
          ProjectActions.projectAutoscaling(this.props.params.projectId, {serviceName: serviceName, scalingType: type});
      } else {
          ProjectActions.projectRemoveAutoscaling(this.props.params.projectId);
      }
  }
  
  printPorts(ports){
      let ret = "";
      if(ports){
          ports.forEach( (value) => {
              ret += ' '+ value.portHost + ' => '+ value.portCont;
          });
      }
      return ret;
  }
  
  handleChangeScale(event){
      var data = this.state.mapServices ? this.state.mapServices : {};
      data[event.target.id] = {
         serviceName: event.target.id,
         number: parseInt(event.target.value)
      };
      
      this.setState({
          mapServices: data,
          services : _.values(data)
        });
  }
  
  render() {
      var services = [];
      if(this.state && this.state.project){
          services = this.state.project.services;
      }
      return (
        <div>
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
                    <div key={value.serviceName}>
                        <h2>{value.serviceName}</h2>
                        <div>
                            Scale : &nbsp;
                            <input type="number" id={value.serviceName} onChange={ (event) => {this.handleChangeScale(event)}}></input>
                            <input type="checkbox" onChange={ (event) => {this.autoscaling("auto", value.serviceName,event)}}>Autoscaling</input>
                            <input type="checkbox" onChange={ (event) => {this.autoscaling("node", value.serviceName, event)}}>Nodescaling</input>
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
                                {value.containers ? value.containers.map(value => (
                                    <TableRow key={value.containerName}>
                                        <TableRowColumn>{value.containerName}</TableRowColumn>
                                        <TableRowColumn>{value.isRunning ? "true" : "false"}</TableRowColumn>
                                        <TableRowColumn>{this.printPorts(value.port)}</TableRowColumn>
                                    </TableRow>
                                )) : ""}
                            </TableBody>
                        </Table>
                    </div>
                ))}
            </div>
            <FloatingActionButton mini={true} style={{marginTop: 10}} 
                onClick={() => {browserHistory.push('/project');}}>
                <ContentBack />
            </FloatingActionButton>
        </div>
      ); 
  }
}
