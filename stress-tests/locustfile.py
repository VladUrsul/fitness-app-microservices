from locust import HttpUser, task, between
import random
import datetime

class FitnessAppUser(HttpUser):
    wait_time = between(1, 1) 

    def on_start(self):
        """Run when a simulated user starts"""
        self.user_id = None
        self.workout_id = None
        self.session_id = None
        self.headers = {
            "Authorization": "Bearer dummy_token"
        }

    @task(2)
    def create_user(self):
        payload = {
            "name": f"User_{random.randint(1, 1000000)}",
            "email": f"user_{random.randint(1, 1000000)}@example.com"
        }
        response = self.client.post("/api/users", json=payload, headers=self.headers, name="/api/users")
        if response.status_code == 200:
            self.user_id = response.json().get("id")

    @task(3)
    def create_workout(self):
        if not self.user_id:
            return
        payload = {
            "title": f"Workout_{random.randint(1, 100)}",
            "duration_minutes": random.randint(20, 90),
            "user_id": self.user_id
        }
        response = self.client.post("/api/workouts", json=payload, headers=self.headers, name="/api/workouts")
        if response.status_code == 200:
            self.workout_id = response.json().get("id")

    @task(4)
    def create_session(self):
        if not self.workout_id:
            return
        now = datetime.datetime.utcnow()
        payload = {
            "workout_id": self.workout_id,
            "started_at": now.isoformat() + "Z",
            "finished_at": (now + datetime.timedelta(minutes=random.randint(20, 90))).isoformat() + "Z"
        }
        response = self.client.post("/api/sessions", json=payload, headers=self.headers, name="/api/sessions")
        if response.status_code == 200:
            self.session_id = response.json().get("id")

    @task(1)
    def get_user(self):
        if self.user_id:
            self.client.get(f"/api/users/{self.user_id}", headers=self.headers, name="/api/users/:id")

    @task(2)
    def get_workout(self):
        if self.workout_id:
            self.client.get(f"/api/workouts/{self.workout_id}", headers=self.headers, name="/api/workouts/:id")

    @task(3)
    def get_session(self):
        if self.session_id:
            self.client.get(f"/api/sessions/{self.session_id}", headers=self.headers, name="/api/sessions/:id")
