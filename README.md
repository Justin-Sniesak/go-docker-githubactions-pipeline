# Go + Docker + GitHub Actions CI/CD Pipeline

A lightweight, end-to-end demo showing how to build, containerize, and automate a Go application with **Docker** and **GitHub Actions**, integrated with **Google Cloud Artifact Registry**.

All commands should be run using bash.

---

## 🚀 Project Overview
This project demonstrates a complete **CI/CD workflow**:

1. **Go app** — simple program that outputs your name, date, and time.  
2. **Dockerized container** — validated locally with `docker ps -a`.  
3. **Artifact Registry integration** — authenticated and pushed via `gcloud`.

### 🕒 Note on Timezone

Please note, time returned is in UTC. If running locally, will default to PDT, this can easily be changed to the timezone of your choice if desired in the code.

---
## Repository Structure
```
├── .github/
    └── workflows/
        ├── README.md
        └── go-dockerimage-pipeline.yaml
├── Deployment
├── Docs
    └── OPERATIONS_LOG.md
├── Dockerfile
├── README.md
└── utc-time-date.go

```
---

## 🧠 What This Shows
- Real-world **CI/CD pipeline** for a containerized Go app.  
- Hands-on use of **GitHub Actions**, **Docker**, and **Google Cloud**.  
- Clean, reproducible workflow for recruiters or hiring managers to verify.

---

## 🧩 Stack
| Layer | Tool |
|-------|------|
| Language | Go |
| Containerization | Docker |
| CI/CD | GitHub Actions |
| Cloud | Google Cloud Platform (Artifact Registry) |

---

## 🧰 Prerequisites
- [Docker](https://docs.docker.com/get-docker/)
- [Go](https://go.dev/doc/install)
- [Google Cloud SDK](https://cloud.google.com/sdk/docs/install)
- A Google Cloud project with an [Artifact Registry](https://cloud.google.com/artifact-registry)

---

## ⚙️ Reproduction Steps

### 1️⃣ Clone the repo
```git clone https://github.com/justin-sniesak/go-docker-githubactions-pipeline.git```

### 2️⃣ Build and run locally
```docker build -t gotime:latest .```

```docker run gotime```

Verify output:

Hello, the date is 10/11/2025, and the time is 07:45 PM.

### 3️⃣ Authenticate with Google Cloud

```gcloud auth configure-docker```

### 4️⃣ Push image to Artifact Registry

```docker tag gotime:latest us-west1-docker.pkg.dev/<your-project>/<repo-name>/gotime:latest```

```docker push us-west1-docker.pkg.dev/<your-project>/<repo-name>/gotime:latest```

### 5️⃣ GitHub Actions CI/CD

Builds Docker image

Outputs current time and date in the pipeline log

### 🧾 Sample Output (GitHub Actions)

```Run docker run --rm gotime```
  ```docker run --rm gotime```
  ```shell: /usr/bin/bash -e {0}```
```Hello, the date is 10/12/2025, and the time is 06:10 AM.```

### 💡 Lessons Learned

How to link GitHub Actions → Docker | Create, authenticate to and push Docker image to GCP Artifact Registry

CI/CD pipelines don’t have to be complex to be production-grade

Debugging is part of mastery — not failure

### 🛠️ Troubleshooting

All screenshots included in repo.

**Issue 1: GCP Authentication Failure**
- **Problem:** Incorrect account was set as active, blocking Docker image push to Artifact Registry
- **Solution:** Switched to correct active account via `gcloud config set account`

**Issue 2: Docker Build Failure**  
- **Problem:** Dockerfile and Go code were not at root level, causing path errors
- **Solution:** Moved both files to root directory to resolve COPY instruction

**Issue 3: Pipeline Failure**
- **Problem:** Renamed code file but didn't update GitHub Actions YAML reference
- **Solution:** Updated workflow file to match new filename, build completed successfully (RC0)

**Issue 4: Tag and Push to GCP Failure**
- **Problem:** Build is failing when attempting to tag the built docker image prior to pushing to Artifact Repo in GCP.
- **Solution:** Added unit test job, tagged image earlier in build process, broke tag/push job into two different jobs

### 🏁 Author

Justin Sniesak 

Infrastructure Engineer | Cloud | Kubernetes | CI/CD | Go
📍 Seattle, WA
