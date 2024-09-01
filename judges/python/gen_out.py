from submission.runner import runner


def main():
    with open("./submission/testin.txt", "r") as test_in:
        with open("./submission/solution.txt", "w") as solution:
            curr_input = ""
            for line in test_in:
                if line == "\f\n":
                    output = runner(curr_input)
                    solution.write(output)
                    solution.write("\n\f\n")
                    continue
                curr_input += line


if __name__ == "__main__":
    main()
