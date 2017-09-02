#!/bin/bash -xe

python manage.py migrate
cat << END | python manage.py shell
from django.contrib.auth.models import User
User.objects.create_superuser('admin', 'admin@inuits.eu', 'pass')
User.objects.create_user('user', 'user@inuits.eu', 'pass')
END
python manage.py loaddata /tests/data.json
python manage.py runserver 0.0.0.0:8000
