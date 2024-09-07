from submission.runner import runner
import sys
from time import process_time


def main():
    with open(sys.argv[1] + "/testin.txt", "r") as test_in:
        with open(sys.argv[1] + "/solution.txt", "w") as solution:
            curr_input = ""
            start = process_time()
            for line in test_in:
                if line == "\f\n":
                    output = runner(curr_input)
                    solution.write(output)
                    solution.write("\n\f\n")
                    continue
                curr_input += line
            end = process_time()
            solution.write(str(end - start))


if __name__ == "__main__":
    main()
