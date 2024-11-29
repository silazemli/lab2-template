
until nc -z your_service_host your_service_port; do
  echo "Waiting for service..."
  sleep 2
done

# Start your main application
exec your_main_application