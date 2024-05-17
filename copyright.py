from pathlib import Path

copyright = """/*
Copyright © 2024 Amr Ojjeh <amrojjeh@outlook.com>
*/
"""

gofiles = Path(".").glob("**/*.go")
exceptions = ["model/db.go", "model/models.go"]

missing = []
for file in gofiles:
    if str(file) in exceptions:
        continue
    with open(file, "r") as f:
        code = f.read()
        if code.find("Copyright") == -1:
            missing.append(file)
            print(file)

if len(missing) == 0:
    print("All files are copyrighted. Nice!")
    exit()

print(f"There are {len(missing)} files missing copyright")
yn = input("Do you want me to fix them (y/n)? ").lower()

if yn != "y" and yn != "yes":
    exit()

for m in missing:
    with open(m, "r+") as f:
        code = f.read()
        code = f"{copyright}\n{code}"
        f.seek(0, 0)
        f.write(code)
