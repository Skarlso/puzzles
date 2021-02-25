package main

type parser struct {
	stars []string
	dot   bool
	// index string
	is int
	// index pattern
	ip      int
	pattern string
	s       string
	next    byte
}

func (p *parser) scan() bool {
	for {
		if p.is >= len(p.s) {
			// for now, add a special case for the greedy matcher
			// if we still have next set, we didn't finish matching yet
			// but we run out of string to match.
			if p.next != 0 {
				return false
			}

			// okay, dirty dirty hack here...
			// check if we didn't end on x* if we did and we matched
			// so far, we return true.
			if p.ip+1 < len(p.pattern) && len(p.pattern) == p.ip+2 {
				if p.pattern[p.ip+1] == '*' {
					return true
				}
			}

			// if there is still pattern to match and we didn't close on a star
			// return false.
			if !p.star && p.ip < len(p.pattern) {
				return false
			}
			break
		}

		// no more string, but there is still pattern left to match and no star is enabled.
		// if p.ip < len(p.pattern) && p.is >= len(p.s)-1 && !p.star {
		// return false
		// }

		if p.ip >= len(p.pattern) && p.is < len(p.s) {
			// no more pattern left but there is still string
			// left to match
			return false
		}

		pcurr := p.pattern[p.ip]
		curr := p.s[p.is]

		// look ahead for tokens if we haven't reached the end of the string
		// atm the only token we care to look ahead is * because
		// that's the only token that needs a classifier.

		// New way. Scan ahead as far as possible and save all
		// the x* matchers into a list until we find something else.
		for {
			offset := 1
			// scan ahead while we find patterns and safe them.
			if p.ip+offset < len(p.pattern) {
				if p.pattern[p.ip+offset] == '*' {
					// maybe also store the stopper for this greedy match?
					// struct for star match?
					c := p.pattern[p.ip+offset]
					p.stars = append(p.stars, string(c))
				}
			}
		}
		// if p.ip+1 < len(p.pattern) {
		// 	switch p.pattern[p.ip+1] {
		// 	case '*':
		// 		// if star is on, we don't move the pattern index
		// 		// and just keep matching the current character
		// 		// until there is no match. in which case we increase the
		// 		// pattern index by two.
		// 		p.star = true

		// 		// check if we have anything after * so .* known until when to match
		// 		// only set next if current is .
		// 		if p.ip+2 < len(p.pattern) && pcurr == '.' {
		// 			p.next = p.pattern[p.ip+2]
		// 		}
		// 	}
		// }

		switch pcurr {
		case '.':
			// greedy match. We match until either we run out of characters to match
			// or, if this is not the end of the pattern, we find a match
			// for whatever is coming after the star.
			if p.star && p.next == 0 {
				// no more patterns, we can greedily match.
				return true
			} else if p.star && p.next != 0 {
				if p.next == curr {
					p.star = false
					p.ip += 2
					p.is++
					p.next = 0
					continue
				}
				p.is++
				continue
			}

			// we match until we run out of string, or encounter the
			// next which is a stopper for the greedy .*.
			p.ip++
			p.is++
			continue
		default:
			// we didn't match, disable star if it was a star.
			if curr != pcurr {
				if p.star {
					// skip the star token it didn't do anything
					// but don't move the string forward.
					p.star = false
					p.ip += 2

					// if we reached the end of the pattern and the
					// string, we just return true. we are done.
					// if p.ip >= len(p.pattern) {
					// 	return true
					// }

					continue
				}
				return false
			}

			// we matched
			if p.star {
				p.is++
				continue
			}

			p.ip++
			p.is++
		}
	}

	return true
}

func isMatch(s string, p string) bool {
	ps := &parser{
		pattern: p,
		s:       s,
	}

	return ps.scan()
}
