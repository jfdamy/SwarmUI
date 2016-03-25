import Reflux from 'reflux';
import ProjectActions from '../actions/projectActions';

var ProjectStore = Reflux.createStore({

    init() {

        this.listenTo(ProjectActions.listProject, this.listProject);
        this.listenTo(ProjectActions.listProjectSuccess, this.listProjectSuccess);
        this.listenTo(ProjectActions.listProjectFailure, this.listProjectFailure);
        
        this.listenTo(ProjectActions.projectInfo, this.projectInfo);
        this.listenTo(ProjectActions.projectInfoSuccess, this.projectInfoSuccess);
        this.listenTo(ProjectActions.projectInfoFailure, this.projectInfoFailure);
        
        this.listenTo(ProjectActions.projectUp, this.projectUp);
        this.listenTo(ProjectActions.projectUpSuccess, this.projectUpSuccess);
        this.listenTo(ProjectActions.projectUpFailure, this.projectUpFailure);
        
        this.listenTo(ProjectActions.projectStop, this.projectStop);
        this.listenTo(ProjectActions.projectStopSuccess, this.projectStopSuccess);
        this.listenTo(ProjectActions.projectStopFailure, this.projectStopFailure);
        
        this.listenTo(ProjectActions.projectKill, this.projectKill);
        this.listenTo(ProjectActions.projectKillSuccess, this.projectKillSuccess);
        this.listenTo(ProjectActions.projectKillFailure, this.projectKillFailure);
        
        this.listenTo(ProjectActions.projectDelete, this.projectDelete);
        this.listenTo(ProjectActions.projectDeleteSuccess, this.projectDeleteSuccess);
        this.listenTo(ProjectActions.projectDeleteFailure, this.projectDeleteFailure);
        
        this.listenTo(ProjectActions.projectRemove, this.projectRemove);
        this.listenTo(ProjectActions.projectRemoveSuccess, this.projectRemoveSuccess);
        this.listenTo(ProjectActions.projectRemoveFailure, this.projectRemoveFailure);
        
        this.listenTo(ProjectActions.projectScale, this.projectScale);
        this.listenTo(ProjectActions.projectScaleSuccess, this.projectScaleSuccess);
        this.listenTo(ProjectActions.projectScaleFailure, this.projectScaleFailure);
        
        this.listenTo(ProjectActions.createProject, this.createProject);
        this.listenTo(ProjectActions.createProjectSuccess, this.createProjectSuccess);
        this.listenTo(ProjectActions.createProjectFailure, this.createProjectFailure);
        
        this.listenTo(ProjectActions.projectDefinition, this.projectDefinition);
        this.listenTo(ProjectActions.projectDefinitionSuccess, this.projectDefinitionSuccess);
        this.listenTo(ProjectActions.projectDefinitionFailure, this.projectDefinitionFailure);
    },

    listProject(){
        this.trigger({});
    },

    listProjectSuccess(projects) {
        this.trigger({ projects: projects});
    },

    listProjectFailure(error){
        this.trigger({error: error});
    },

    projectInfo(){
        this.trigger({});
    },

    projectInfoSuccess(project) {
        this.trigger({project: project});
    },

    projectInfoFailure(error){
        this.trigger({error: error});
    },

    projectDefinition(){
        this.trigger({});
    },

    projectDefinitionSuccess(project) {
        this.trigger(project);
    },

    projectDefinitionFailure(error){
        this.trigger({error: error});
    },
    
    projectUp(){
        this.trigger({});
    },

    projectUpSuccess(project) {
        this.trigger({refreshProject: true});
    },

    projectUpFailure(error){
        this.trigger({error: error});
    },
    
    projectStop(){
        this.trigger({});
    },

    projectStopSuccess(project) {
        this.trigger({refreshProject: true});
    },

    projectStopFailure(error){
        this.trigger({error: error});
    },
    
    projectKill(){
        this.trigger({});
    },

    projectKillSuccess(project) {
        this.trigger({refreshProject: true});
    },

    projectKillFailure(error){
        this.trigger({error: error});
    },
    
    projectDelete(){
        this.trigger({});
    },

    projectDeleteSuccess(project) {
        this.trigger({refreshProject: true});
        this.trigger({returnToProjects: true});
    },

    projectDeleteFailure(error){
        this.trigger({returnToProjects: true, error: error});
    },
    
    projectRemove(){
        this.trigger({});
    },

    projectRemoveSuccess(project) {
        this.trigger({refreshProject: true});
    },

    projectRemoveFailure(error){
        this.trigger({error: error});
    },
    
    projectScale(){
        this.trigger({});
    },

    projectScaleSuccess(project) {
        this.trigger({refreshProject: true});
    },

    projectScaleFailure(error){
        this.trigger({error: error});
    },
    
    createProject(){
        this.trigger({});
    },

    createProjectSuccess(project) {
        this.trigger({refreshProject: true});
    },

    createProjectFailure(error){
        this.trigger({error: error});
    },
});

export default ProjectStore;