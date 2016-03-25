import Reflux from 'reflux';
import Api from '../utils/api'

var ProjectActions = Reflux.createActions([
    'listProject',
    'listProjectSuccess',
    'listProjectFailure',
    
    'projectInfo',
    'projectInfoSuccess',
    'projectInfoFailure',
    
    'projectDefinition',
    'projectDefinitionSuccess',
    'projectDefinitionFailure',
    
    'projectUp',
    'projectUpSuccess',
    'projectUpFailure',
    
    'projectStop',
    'projectStopSuccess',
    'projectStopFailure',
    
    'projectKill',
    'projectKillSuccess',
    'projectKillFailure',
    
    'projectDelete',
    'projectDeleteSuccess',
    'projectDeleteFailure',
    
    'projectRemove',
    'projectRemoveSuccess',
    'projectRemoveFailure',
    
    'projectScale',
    'projectScaleSuccess',
    'projectScaleFailure',
    
    'createProject',
    'createProjectSuccess',
    'createProjectFailure',
]);

ProjectActions.listProject.preEmit = function(){
    Api.listProject()
        .error(function(err){
            console.log(err);
            ProjectActions.listProjectFailure(err);
        })
        .then((result) => {
            ProjectActions.listProjectSuccess(result);
        });
};

ProjectActions.createProject.preEmit = function(projectId, composeDefinition){
    Api.createProject(projectId, composeDefinition)
        .error(function(err){
            console.log(err);
            ProjectActions.createProjectFailure(err);
        })
        .then((result) => {
            ProjectActions.createProjectSuccess(result);
        });
};

ProjectActions.projectInfo.preEmit = function(projectId){
    Api.projectInfo(projectId)
        .error(function(err){
            console.log(err);
            ProjectActions.projectInfoFailure(err);
        })
        .then((result) => {
            ProjectActions.projectInfoSuccess(result);
        });
};

ProjectActions.projectDefinition.preEmit = function(projectId){
    Api.projectDefinition(projectId)
        .error(function(err){
            console.log(err);
            ProjectActions.projectDefinitionFailure(err);
        })
        .then((result) => {
            ProjectActions.projectDefinitionSuccess(result);
        });
};

ProjectActions.projectUp.preEmit = function(projectId, services = []){
    Api.projectUp(projectId, {ServicesName: services})
        .error(function(err){
            console.log(err);
            ProjectActions.projectUpFailure(err);
        })
        .then((result) => {
            ProjectActions.projectUpSuccess(result);
        });
};

ProjectActions.projectStop.preEmit = function(projectId, services = []){
    Api.projectStop(projectId, {ServicesName: services})
        .error(function(err){
            console.log(err);
            ProjectActions.projectStopFailure(err);
        })
        .then((result) => {
            ProjectActions.projectStopSuccess(result);
        });
};

ProjectActions.projectKill.preEmit = function(projectId, services = []){
    Api.projectKill(projectId, {ServicesName: services})
        .error(function(err){
            console.log(err);
            ProjectActions.projectKillFailure(err);
        })
        .then((result) => {
            ProjectActions.projectKillSuccess(result);
        });
};

ProjectActions.projectDelete.preEmit = function(projectId, services = []){
    Api.projectDelete(projectId, {ServicesName: services})
        .error(function(err){
            console.log(err);
            ProjectActions.projectDeleteFailure(err);
        })
        .then((result) => {
            ProjectActions.projectDeleteSuccess(result);
        });
};

ProjectActions.projectRemove.preEmit = function(projectId){
    Api.projectRemove(projectId)
        .error(function(err){
            console.log(err);
            ProjectActions.projectRemoveFailure(err);
        })
        .then((result) => {
            ProjectActions.projectRemoveSuccess(result);
        });
};

ProjectActions.projectScale.preEmit = function(projectId, servicesScale){
    Api.projectScale(projectId, servicesScale)
        .error(function(err){
            console.log(err);
            ProjectActions.projectScaleFailure(err);
        })
        .then((result) => {
            ProjectActions.projectScaleSuccess(result);
        });
};

export default ProjectActions;