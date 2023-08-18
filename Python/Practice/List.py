if __name__ == "__main__":
  people = ["W", "A"]
  print(people)

  people.append("P")
  print(people)

  peopleCopy = people.copy()
  peopleCopy.append("P")
  print(peopleCopy, people)

  print(peopleCopy.count("P"))

  peopleCopy.extend(people)
  print(peopleCopy)

  print(peopleCopy.index("P"))

  peopleCopy.insert(3, "J")
  print(peopleCopy)

  peopleCopy.pop(4)
  print(peopleCopy)

  peopleCopy.remove("J")
  print(peopleCopy)

  peopleCopy.reverse()
  print(peopleCopy)

  peopleCopy.sort(reverse=True)
  print(peopleCopy)