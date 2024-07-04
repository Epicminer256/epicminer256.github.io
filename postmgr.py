import uuid
import json

print("Loading file...")
posts = json.load(open("blog.json", "r"))
print("Done!\n")

def writePosts():
    json.dump(posts, open("blog.json", "w"), indent=4)

while True:
    print("Type a letter or use \"H\" for help:")
    command = input("==> ")
    try:
        command = command[0]
    except IndexError:
        print("Please type a command")
        continue
    
    if command == "h":
        print("[H]elp")
        print("[N]ew post")
        print("[O]pen post")
        print("[Q]uit")
        print("\nYou only type one letter into the terminal\n")
        continue
    
    if command == "q":
        print("Exiting...")
        break
    
    if command == "n":
        sections=[]
        print("Enter a title (leave empty to quit)")
        title=input(">")
        if title=="": continue
        while True:
            print("Enter a section name (leave empty to finalyze post)")
            header=input(">")
            if header=="": break
            print("Enter the section's contents (leave empty to cancel making this section)")
            content=input(">")
            if title=="": continue
            sections.append({"header": header, "content": content})
        
        rawPost = {
            "UUID": str(uuid.uuid4()),
            "title": title,
            "content": sections
        }
        
        rawPostJSON = json.dumps(rawPost, indent=4)
        print(rawPostJSON)
        del rawPostJSON

        print("\nDoes everything look good to post?")
        while True:
            confirmation=input("[y] or [n]>")
            try:
                confirmation = confirmation[0].lower()
            except IndexError:
                print("Please make a answer")
                continue
            if confirmation == "y":
                posts.append(rawPost)
                writePosts()
                break

            if confirmation == "n":
                break
            
            print("You entered a invalid response")
        continue
    
    if command == "o":
        if len(posts) == 0:
            print("There are no posts")
            continue

        for i in range(0, len(posts)):
            index=i+1
            print("["+str(index)+"] "+posts[i]["title"])
        while True:
            print("Select a post to open (Leave empty to quit)")
            response = input(">")
            if response == "": break
            try:
                response = int(response)
            except ValueError:
                print("Invalid response!")
                continue
            try:
                post = posts[response-1]
                print("\n\n===== "+post["title"]+" =====")
                for section in post['content']:
                    print("\n--- "+section['heading']+"---")
                    print(section['content'])
                print("\n")
                break
            except IndexError:
                print("Post not found!")
        continue
    
    print("Command not found")