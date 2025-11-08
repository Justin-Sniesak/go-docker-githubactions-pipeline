## Troubleshooting Callout
Artifacts of smoke testing pipeline build, includes troubleshooting.

## Pipeline Build
**Summary:** Write code -> build docker image -> test code locally -> create repo -> test and debug pipeline -> clone repo and test locally.

- 2025-10-11 Build docker image and run container.
  ![Dep1-1](../Deployment/Dep1-1.jpg)
- 2025-10-11 Validate container ran and returned RC0.
  ![Dep1-2](../Deployment/Dep1-2.jpg)
- 2025-10-11 Push dockerfile to GCP Artifact Repository.
  ![Dep1-3](../Deployment/Dep1-3.jpg)
- 2025-10-11 Authorize Docker to connect to Artifact Repo using gcloud shell.
  ![Dep1-4](../Deployment/Dep1-4.jpg)
- 2025-11-11 TROUBLESHOOT: The incorrect account was set as active, cuasing authentication to fail when trying to push the docker image to GCP Artifact Repo. Corrected by changing and validating the correct active account.
  ![Dep1-5](../Deployment/Dep1-5.jpg)
- 2025-10-11 Validate dockerfile is now in Artifact Repo in GCP.
  ![Dep1-6](../Deployment/Dep1-6.jpg)
- 2025-10-11 TROUBLESHOOT: Dockerfile and Go code where not both in root, which was causing the build to fail.
  ![Dep1-7](../Deployment/Dep1-7.jpg)
- 2025-10-11 TROUBLESHOOT: Updated the name of the code used to build the Docker image, did not update the build yaml or the code itself, which was causing the build to fail. Once everything was uniform build ran to completion returning an RC0.
  ![Dep1-8](../Deployment/Dep1-8.jpg)
- 2025-10-12 TROUBLESHOOT: Job failing due to the source image not being properly tagged prior to pushing the docker image from GitHub Actions to GCP Artifact Repo. Corrected tagging the image earlier in the build process, adding a unit test job, and breaking up the tag and push job from one to two jobs. Final smoke test performed by validating image updated in Artifact Repo.
  ![Dep1-9](../Deployment/Dep1-9.jpg)
  ![Dep1-10](../Deployment/Dep1-10.jpg)
  ![Dep1-11](../Deployment/Dep1-11.jpg)
