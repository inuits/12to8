#!/bin/bash -xe

if [[ -n ${FIXTURE} ]]
then
    python manage.py migrate  --configuration Prod
    cat << END | python manage.py shell  --configuration Prod
from django.contrib.auth.models import User
User.objects.create_superuser('admin', 'admin@inuits.eu', 'pass')
User.objects.create_user('user', 'user@inuits.eu', 'pass')
END
    exec python manage.py loaddata /tests/fixtures/${FIXTURE}.json --configuration Prod
else
    exec python manage.py runserver 0.0.0.0:8000 --configuration Prod
fi
