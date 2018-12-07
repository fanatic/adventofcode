#!/usr/bin/env python3

""" Advent of code """

import re
from collections import namedtuple


def part_1(fh):
    graph = {}
    all_steps = set()
    for line in fh:
        m = re.match(
            r"Step (.) must be finished before step (.) can begin.", line)
        all_steps.add(m.group(1))
        all_steps.add(m.group(2))
        try:
            graph[m.group(2)].add(m.group(1))
        except KeyError:
            graph[m.group(2)] = {m.group(1)}

    return ''.join(find_path(all_steps, graph))


def find_path(all_steps, graph, path=[]):
    print(':', all_steps, graph, path)
    # Find the step with no prereqs
    next_step = find_next_step(all_steps, graph)

    # If none left, we're done
    if next_step == None:
        return path

    # Otherwise, recurse
    return find_path(all_steps - {next_step}, graph, path + [next_step])


def find_next_step(all_steps, graph, ignore=set()):
    possible = set()
    for s in sorted(all_steps):
        try:
            if len(graph[s] & all_steps) == 0:
                possible.add(s)
        except KeyError:
            possible.add(s)
    next_steps = (possible - ignore)
    if len(next_steps) == 0:
        return None

    return sorted(next_steps)[0]


Worker = namedtuple('Worker', ['step', 'time_spent'])
step_duration = 60
num_workers = 5


def part_2(fh):
    graph = {}
    all_steps = set()
    for line in fh:
        m = re.match(
            r"Step (.) must be finished before step (.) can begin.", line)
        all_steps.add(m.group(1))
        all_steps.add(m.group(2))
        try:
            graph[m.group(2)].add(m.group(1))
        except KeyError:
            graph[m.group(2)] = {m.group(1)}

    workers = {}
    for i in range(num_workers):
        workers[i] = Worker('', 0)
    done = set()
    print("Second", "Workers", "Done")
    second = 0
    while True:
        done |= per_second(graph, all_steps - done, workers)
        print(second, workers, done)

        if len(done) == len(all_steps):
            break
        second += 1

    return second


def per_second(graph, all_steps, workers):
    finished_tasks = set()
    for i in workers:
        # Find finished tasks
        if workers[i].step != '':
            task_time = step_duration + ord(workers[i].step)-64
            if workers[i].time_spent == task_time:
                finished_tasks.add(workers[i].step)
                all_steps = all_steps - {workers[i].step}
                workers[i] = Worker('', 0)

        # Assign Work
        if workers[i].step == '':
            # Find next task (and our exit condition)
            next_step = find_next_step(
                all_steps, graph, steps_in_progress(workers))
            #print('find_next-step', all_steps, next_step)
            if next_step != None:
                workers[i] = Worker(next_step, 0)

        # Assign Hours
        if workers[i].step != '':
            workers[i] = Worker(workers[i].step, workers[i].time_spent+1)
    return finished_tasks


def steps_in_progress(workers):
    steps = set()
    for i in workers:
        if workers[i].step != '':
            steps.add(workers[i].step)
    return steps


def main():
    "Main Entrypoint"
    print('Part 1 is: ', part_1(open('input_7', 'r')))
    print('Part 2 is: ', part_2(open('input_7', 'r'))-1)


if __name__ == "__main__":
    main()
