read -p "Enter the name of the migration (e.g., create_users_table): " migration_name

# Generate timestamp
timestamp=$(date +%s)

# Create migration files
touch database/migrations/"${timestamp}_${migration_name}.up.sql"
touch database/migrations/"${timestamp}_${migration_name}.down.sql"

echo "Migration files created:"
echo "${timestamp}_${migration_name}.up.sql"
echo "${timestamp}_${migration_name}.down.sql"