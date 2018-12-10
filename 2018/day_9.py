#!/usr/bin/env python3

""" Advent of code """


def part_1(players, points):
    player_scores = [0] * players
    marbles = [0]
    position = 0
    #print_marbles(marbles, position, '-')

    for current_marble in range(1, points+1):
        player = current_marble % players

        if current_marble % 23 == 0:
            position = (position - 7) % len(marbles)
            player_scores[player-1] += current_marble + marbles[position+1]
            # print(player, player_scores[player-1])
            del marbles[position+1]
        else:
            position = (position + 2) % len(marbles)
            marbles.insert(position+1, current_marble)

        # print_marbles(marbles, position+1, player)

    # print('scores', player_scores)
    return max(player_scores)


def print_marbles(marbles, position, player):
    if len(marbles) > 100:
        return
    print('[%s]' % player, end=' ')
    for i in range(len(marbles)):
        if position == i:
            print('(%d) ' % marbles[i], end='')
        else:
            print('%d ' % marbles[i], end='')
    print()


def main():
    "Main Entrypoint"
    print('Part 1 is: ', part_1(419, 71052))
    print('Part 2 is: ', part_1(419, 71052*100))


if __name__ == "__main__":
    main()
