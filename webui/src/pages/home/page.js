import React from "react";
import styles from "./style.css";
import ProjectActions from '../../actions/projectActions';
import ProjectStore from '../../stores/projectStore';
import { browserHistory, Link } from 'react-router'

import Table from 'material-ui/lib/table/table';
import TableHeaderColumn from 'material-ui/lib/table/table-header-column';
import TableRow from 'material-ui/lib/table/table-row';
import TableHeader from 'material-ui/lib/table/table-header';
import TableRowColumn from 'material-ui/lib/table/table-row-column';
import TableBody from 'material-ui/lib/table/table-body';
import FloatingActionButton from 'material-ui/lib/floating-action-button';
import ContentAdd from 'material-ui/lib/svg-icons/content/add';



export default class HomePage extends React.Component {
    
  componentDidMount(){
      this.unsubscribe = ProjectStore.listen((state) => {this.onListChange(state);});
      ProjectActions.listProject();
  }
  
  componentWillUnmount(){
      this.unsubscribe();
  }
  
  onListChange(state){
      this.setState(state);
  }
  
  render() {
    var projectsList = [];
    if(this.state && this.state.projects){
        projectsList = this.state.projects;
    }
    const style = {
        marginTop: 20,
        marginLeft: "90%"
    };
    return (
      <div className={styles.content}>
        <h1>List of project</h1>
        <Table>
            <TableHeader>
            <TableRow>
                <TableHeaderColumn>Name</TableHeaderColumn>
                <TableHeaderColumn>Actions</TableHeaderColumn>
            </TableRow>
            </TableHeader>
            <TableBody>
                {projectsList.map(value => (
                    <TableRow key={value}>
                        <TableRowColumn>{value}</TableRowColumn>
                        <TableRowColumn>
                            <Link to={`/project/${value}`}>Show</Link>
                            <br />
                            <Link to={`/edit/${value}`}>Edit</Link>
                        </TableRowColumn>
                    </TableRow>
                ))}
            </TableBody>
        </Table>
        <FloatingActionButton style={style} onClick={() => {browserHistory.push('/create/project');}}>
            <ContentAdd />
        </FloatingActionButton>
      </div>
    );
  }
}
