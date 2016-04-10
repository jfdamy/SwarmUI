var request = require('superagent');
import { Promise } from 'bluebird';

class Api {
    constructor(){
        this.httpPrefix ='/api/v1';
    }

    listProject() {
        return new Promise((resolve, reject) => {
            request.get(this.httpPrefix + '/project')
                .set('Accept', 'application/json')
                .end(function (err, res) {
                    if(err){
                        reject(err);
                    } else {
                        resolve(res.body);
                    }
                });
        });
    }
    
    projectInfo(projectId) {
        return new Promise((resolve, reject) => {
            request.get(this.httpPrefix + '/project/'+projectId)
                .set('Accept', 'application/json')
                .end(function (err, res) {
                    if(err){
                        reject(err);
                    } else {
                        resolve(res.body);
                    }
                });
        });
    }
    
    projectDefinition(projectId) {
        return new Promise((resolve, reject) => {
            request.get(this.httpPrefix + '/project/'+projectId+'/definition')
                .set('Accept', 'application/json')
                .end(function (err, res) {
                    if(err){
                        reject(err);
                    } else {
                        resolve(res.body);
                    }
                });
        });
    }
    
    createProject(projectId, definition) {
        return new Promise((resolve, reject) => {
            request.post(this.httpPrefix + '/project/'+projectId)
                .set('Accept', 'application/yaml')
                .send(definition)
                .end(function (err, res) {
                    if(err){
                        reject(err);
                    } else {
                        resolve(res.body);
                    }
                });
        });
    }
    
    projectUp(projectId, services) {
        return new Promise((resolve, reject) => {
            request.post(this.httpPrefix + '/project/'+projectId+'/up')
                .set('Accept', 'application/json')
                .send(services)
                .end(function (err, res) {
                    if(err){
                        reject(err);
                    } else {
                        resolve(res.body);
                    }
                });
        });
    }
    
    projectStop(projectId, services) {
        return new Promise((resolve, reject) => {
            request.post(this.httpPrefix + '/project/'+projectId+'/stop')
                .set('Accept', 'application/json')
                .send(services)
                .end(function (err, res) {
                    if(err){
                        reject(err);
                    } else {
                        resolve(res.body);
                    }
                });
        });
    }
    
    projectKill(projectId, services) {
        return new Promise((resolve, reject) => {
            request.post(this.httpPrefix + '/project/'+projectId+'/kill')
                .set('Accept', 'application/json')
                .send(services)
                .end(function (err, res) {
                    if(err){
                        reject(err);
                    } else {
                        resolve(res.body);
                    }
                });
        });
    }
    
    projectDelete(projectId, services) {
        return new Promise((resolve, reject) => {
            request.post(this.httpPrefix + '/project/'+projectId+'/delete')
                .set('Accept', 'application/json')
                .send(services)
                .end(function (err, res) {
                    if(err){
                        reject(err);
                    } else {
                        resolve(res.body);
                    }
                });
        });
    }
    
    projectRemove(projectId) {
        return new Promise((resolve, reject) => {
            request.post(this.httpPrefix + '/project/'+projectId+'/remove')
                .set('Accept', 'application/json')
                .end(function (err, res) {
                    if(err){
                        reject(err);
                    } else {
                        resolve(res.body);
                    }
                });
        });
    }
    
    projectScale(projectId, servicesScale) {
        return new Promise((resolve, reject) => {
            request.post(this.httpPrefix + '/project/'+projectId+'/scale')
                .set('Accept', 'application/json')
                .send(servicesScale)
                .end(function (err, res) {
                    if(err){
                        reject(err);
                    } else {
                        resolve(res.body);
                    }
                });
        });
    }
    
    projectAutoscaling(projectId, servicesScale) {
        return new Promise((resolve, reject) => {
            request.post(this.httpPrefix + '/project/'+projectId+'/autoscaling')
                .set('Accept', 'application/json')
                .send(servicesScale)
                .end(function (err, res) {
                    if(err){
                        reject(err);
                    } else {
                        resolve(res.body);
                    }
                });
        });
    }
    
    projectRemoveAutoscaling(projectId) {
        return new Promise((resolve, reject) => {
            request.delete(this.httpPrefix + '/project/'+projectId+'/autoscaling')
                .set('Accept', 'application/json')
                .end(function (err, res) {
                    if(err){
                        reject(err);
                    } else {
                        resolve(res.body);
                    }
                });
        });
    }
}

export default new Api();