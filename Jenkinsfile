pipeline {
    agent any
    tools {
        go 'go1.19'
    }
    environment {
        DOCKERHUB_CREDENTIALS = credentials('sj-dockerhub')
     }
    stages {
        stage('Build Go Project') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: '*/main']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/SourabhJoshi209/devops-integration.git']]])
                
            }
        }
        stage('Build Docker Image') {
            steps {
               script {
                bat 'docker build -t sourabhjoshi209/album-api .'
                
               } 
            }
            
        }
        stage('Push image') {

			steps {
				    bat 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
            bat 'docker push sourabhjoshi209/album-api'
			    }
		   }
    }
}
